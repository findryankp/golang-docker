# Deployment

## GCP
* Buat Instance (Compute Engine): https://cloud.google.com/compute/docs/instances/create-start-instance
* Generate SSH untuk connect ke Instance GCP: https://cloud.google.com/compute/docs/connect/create-ssh-keys#windows-10-or-later
* Menambahkan SSH ke Instance: https://cloud.google.com/compute/docs/connect/add-ssh-keys
* Setup Firewall Rule: https://cloud.google.com/vpc/docs/using-firewalls

Note saat melakukan generate SSH di windows, jika nama user windows nya ada spasi, maka tambahkan tanda "" di url tempat penyimpanan file ssh nya.
contoh "c:/Users/Fakhry Firdaus/.ssh/sshgcpalta"

## Connect to Compute Engine / EC2
```bash
ssh -i </directory/namafilessh> <username-server>@<public-ipv4>

# example:
[GCP] ssh -i ~/.ssh/sshgcp fakhry@18.0.1.2

[AWS] ssh -i ~/.ssh/file.pem ubuntu@18.0.1.2
```
Note:
* Untuk connect ke GCP, kita perlu generate ssh key terlebih dahulu, dan menambahkan ssh key public nya ke setup instance GCP.
* Untuk connect ke AWS, kita dapat menggunakan file .pem yang didapat saat awal pembuatan Instance untuk connect ke Instance AWS. 
* untuk nama username di GCP, biasanya sesuai dengan nama alamat email kita. misal fakhry@gmail.com
maka username instance kita adalah fakhry.
* Untuk AWS, jika anda menggunakan ubuntu. maka username nya adalah ubuntu.