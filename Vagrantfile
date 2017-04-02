# -*- mode: ruby -*-
# vi: set ft=ruby :

# VM Customized Settings
$CPUS              = 1
$MEMORY            = 512


# Setup systemd service file
# Creates and enable systemd service
$setup_systemd = <<SCRIPT
cat > /etc/systemd/system/selenium_hub.service <<-EOF
[Unit]
Description=Selenium server hub

[Service]
TimeoutStartSec=10
RestartSec=10
ExecStart=/bin/bash -c 'java -jar /home/vagrant/selenium-server-standalone-3.3.1.jar -role hub'
WorkingDirectory=/vagrant
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOF


cat > /etc/systemd/system/selenium_node.service <<-EOF
[Unit]
Description=Selenium server node
After=selenium_hub.service
Wants=selenium_hub.service

[Service]
TimeoutStartSec=10
RestartSec=10
ExecStart=/bin/bash -c 'java -jar /home/vagrant/selenium-server-standalone-3.3.1.jar -role node -hub http://localhost:4444/grid/register'
WorkingDirectory=/vagrant
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOF


systemctl enable selenium_hub.service
systemctl enable selenium_node.service
systemctl daemon-reload
systemctl start selenium_hub.service
systemctl start selenium_node.service
SCRIPT

# Downloads selenium server and webdrivers
$download_selenium = <<SCRIPT
echo "working directory"
echo "`pwd`"

if [ ! -f "selenium-server-standalone-3.3.1.jar" ]; then
    wget http://selenium-release.storage.googleapis.com/3.3/selenium-server-standalone-3.3.1.jar
fi

# Get browser drivers
if [ ! -f "chromedriver" ]; then
    wget https://chromedriver.storage.googleapis.com/2.28/chromedriver_linux64.zip
    unzip chromedriver*.zip
    rm chromedriver*.zip
fi

if [ ! -f "geckodriver" ]; then
    wget https://github.com/mozilla/geckodriver/releases/download/v0.15.0/geckodriver-v0.15.0-linux64.tar.gz
    tar -xvf geckodriver*.tar.gz
    rm geckodriver*.tar.gz
fi
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

	# Assign this VM to a host only network IP, allowing yous to access it
	# via the IP.
	#config.vm.network "private_network", ip: "172.20.0.10", netmask: "255.240.0.0", :mac => "08002719318B"

	# Forward a port from the guest to the host, which allows for outside
	# computers to access the VM, whereas host only networking does not.
	config.vm.network :forwarded_port, guest: 4444, host: 4444

	# Share an additional folder to the guest VM. The first argument is
	# an identifier, the second is the path on the guest to mount the
	# folder, and the third is the path on the host to the actual folder.
	#config.vm.synced_folder ".", "/vagrant", type: "virtualbox"
	#config.vm.synced_folder ".", "/vagrant", type: "rsync"

	# Enable provisioning with a shell script
	config.vm.provision "shell", inline: 'echo "deb http://http.debian.net/debian jessie-backports main" > /etc/apt/sources.list.d/jessie_backports.list'
	config.vm.provision "shell", inline: 'aptitude update'
	#config.vm.provision "shell", inline: 'aptitude upgrade'
	config.vm.provision "shell", inline: 'aptitude -yy install unzip'
	config.vm.provision "shell", inline: 'aptitude -yy install default-jre'
	config.vm.provision "shell", inline: 'aptitude -yy install default-jdk'
	#config.vm.provision "shell", inline: 'aptitude -yy install openjdk-8-jre-headless'
	config.vm.provision "shell", inline: 'aptitude -yy install -t jessie-backports openjdk-8-jre-headless ca-certificates-java'
	config.vm.provision "shell", inline: 'aptitude -yy install openjdk-8-jdk'
	config.vm.provision "shell", inline: '/usr/sbin/update-java-alternatives -s java-1.8.0-openjdk-amd64'
	config.vm.provision "shell", inline: $download_selenium
	config.vm.provision "shell", inline: $setup_systemd
	config.vm.provision "shell", run: "always", inline: "systemctl restart selenium_hub.service"
	config.vm.provision "shell", run: "always", inline: "systemctl restart selenium_node.service"

	config.vm.provider "virtualbox" do |v|
		v.memory = $MEMORY
		v.cpus = $CPUS
	end
end



#echo "deb http://ftp.de.debian.org/debian jessie-backports main" > /etc/apt/sources.list.d/jessie_backports
