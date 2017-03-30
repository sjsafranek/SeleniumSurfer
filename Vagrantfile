# -*- mode: ruby -*-
# vi: set ft=ruby :

# VM Customized Settings
$CPUS              = 1
$MEMORY            = 512


# Setup systemd service file
# Creates and enable systemd service
$setup_systemd = <<SCRIPT
cat > /etc/systemd/system/selenium_server.service <<-EOF
[Unit]
Description=The Selenium Server

[Service]
TimeoutStartSec=10
RestartSec=10
ExecStart=/vagrant/start_selenium_server.sh
WorkingDirectory=/vagrant
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOF

systemctl enable selenium_server.service
systemctl daemon-reload
systemctl start selenium_server.service
SCRIPT


# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

#Vagrant::Config.run do |config|
Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
	# Base box to build off, and download URL for when it doesn't exist on the user's system already
	#config.vm.box = "ubuntu/trusty64"
	#config.vm.box = "debian/jessie64"
	# "debian/jessie64" has a bug with `synced_folder` impacting guest and host sharing of `/vagrant`
	config.vm.box = "debian/contrib-jessie64"
	
	# Boot with a GUI so you can see the screen. (Default is headless)
	# config.vm.boot_mode = :gui
	
	# Assign this VM to a host only network IP, allowing you to access it
	# via the IP.
	#config.vm.network "private_network", ip: "172.20.0.10", netmask: "255.240.0.0", :mac => "08002719318B"
	
	# Forward a port from the guest to the host, which allows for outside
	# computers to access the VM, whereas host only networking does not.
	config.vm.network :forwarded_port, guest: 8080, host: 8080
	
	# Share an additional folder to the guest VM. The first argument is
	# an identifier, the second is the path on the guest to mount the
	# folder, and the third is the path on the host to the actual folder.	
	#config.vm.synced_folder ".", "/vagrant", type: "virtualbox"
	#config.vm.synced_folder ".", "/vagrant", type: "rsync"

	# Enable provisioning with a shell script.
	# sudo apt-add-repository ppa:ubuntugis/ubuntugis-unstable
	config.vm.provision "shell", inline: 'aptitude update'
	config.vm.provision "shell", inline: 'aptitude -yy install openjdk-8-jre-headless '
	config.vm.provision "shell", inline: 'aptitude -yy install openjdk-8-jdk'
	config.vm.provision "shell", inline: '/usr/sbin/update-java-alternatives -s java-1.8.0-openjdk-amd64'
	config.vm.provision "shell", inline: download_selenium_linux.sh
	config.vm.provision "shell", inline: $setup_systemd
	config.vm.provision "shell", run: "always", inline: "systemctl restart selenium_server.service"

	config.vm.provider "virtualbox" do |v|
		v.memory = $MEMORY
		v.cpus = $CPUS
	end
end
