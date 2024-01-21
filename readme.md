# Simple Ansible Playbook: Configure and Run Go Web Server

## Overview
This Ansible playbook automates the deployment and configuration of a Go web server on remote servers. It performs tasks such as creating directories, copying files, and configuring services to set up a functional Go web application.

## Prerequisites
- Ansible installed on the control machine.
- Access to remote servers via SSH.
- Proper SSH credentials (e.g., private key) for connecting to remote servers.

## Usage
1. Clone this repository:

    ```bash
    git clone https://github.com/your-username/your-repo.git
    cd your-repo
    ```

2. Update the Ansible inventory (hosts file) with your server details:

    ```bash
    [web_servers]
    server1 ansible_user=your-ssh-user ansible_ssh_private_key_file=your-ssh-key.pem
    ```
3. Modify playbook variables and configurations (if needed) in playbook.yaml.

4. Run the Ansible playbook:

    ```bash
    ansible-playbook -i hosts playbook.yaml
    ```
## Playbook Structure
1. Create Destination Directory: Ensure the existence of the `/home/app-ansible directory`.
2. Remove Existing goapp File: Delete any existing goapp file in `/home/app-ansible`.
3. Copy New Binary File: Copy the latest binary file to `/home/app-ansible/goapp`.
4. Copy Nginx Configuration: Place the go.conf Nginx configuration file in `/etc/nginx/conf.d/`.
5. Reload systemd (daemon-reload): Execute `systemctl daemon-reload` to reload systemd configurations.
6. Reload Nginx: Restart the Nginx service to apply the new configuration.
7. Configure systemd Service: Use a template to configure the goapp systemd service.
8. Restart goapp Service: Restart the goapp service using sudo service goapp restart.

## Notes
- This playbook assumes the target servers are running systemd and Nginx.
- Customize the playbook according to your specific environment and requirements.