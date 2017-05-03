sudo systemctl restart network
sudo yum -y install epel-release
sudo yum -y install htop vim
sudo yum -y group install "Development Tools"
     
# Set Git User/Email
sudo -H -u vagrant bash -c "git config --global user.name \'#{GITUSER}\'"
sudo -H -u vagrant bash -c "git config --global user.email \'#{GITUSEREMAIL}\'"
     
# Golang
wget -q https://storage.googleapis.com/golang/go1.8.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.8.1.linux-amd64.tar.gz
     
# Set Golang env vars
sed -ie '$aexport GOPATH=\/home\/vagrant\/go' /home/vagrant/.bashrc
sed -ie '$aexport PATH=\$PATH:\/usr\/local\/go\/bin' /home/vagrant/.bashrc
sed -ie '$aexport PATH=\$PATH:\$GOPATH\/bin' /home/vagrant/.bashrc
source /home/vagrant/.bashrc
     
# Make Golang dir structure
sudo mkdir -p /home/vagrant/go/bin
sudo mkdir -p /home/vagrant/go/pkg
sudo mkdir -p /home/vagrant/go/src
sudo chown -R vagrant:vagrant /home/vagrant/go
     
# Glide
wget -q https://github.com/Masterminds/glide/releases/download/v0.12.3/glide-v0.12.3-linux-amd64.tar.gz
tar -C /home/vagrant/go/bin/ -xvf glide-v0.12.3-linux-amd64.tar.gz
sudo mv /home/vagrant/go/bin/linux-amd64/glide /home/vagrant/go/bin/.