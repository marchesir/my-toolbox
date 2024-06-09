#!/bin/bash 

# https://www.gnu.org/software/bash/manual/html_node/The-Set-Builtin.html
set -uo pipefail

export VM_NAME=vscode-remote
export CLOUDINIT_NAME=$VM_NAME

generate_sshkey_pair()
{
    ssh-keygen -q -t ed25519 -a 150 -N '' -f $HOME/.ssh/id_ed25519 <<<y >/dev/null 2>&1
    echo $(cat $HOME/.ssh/id_ed25519.pub)
}

generate_sshconfig()
{
cat > $HOME/.ssh/config <<EOF
Host $vm_ip 
    HostName $vm_ip 
    User ubuntu 
EOF
}

generate_vm_cloudinit() {
cat > $HOME/${CLOUDINIT_NAME}.yaml <<EOF
#cloud-config
package_update: true
package_upgrade: true
ssh_authorized_keys:
  - $sshkey_public
runcmd:
  - curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh && sudo usermod -aG docker ubuntu
EOF
}

is_vm_active() 
{
    multipass info $VM_NAME >/dev/null 2>&1
    [[ $? == 0 ]] && echo "true" || echo "false"
}

get_vm_ip() {
    echo $(multipass info $VM_NAME | grep IPv4 | awk '{print $2}')
}

# 1. check if vm is active, if so go direct to 5. 
if [[ $(is_vm_active) == "false" ]]; then
    # 2. generate sshkey pair id needed.
    [[ ! -f $HOME/.ssh/id_ed25519.pub ]] && sshkey_public=$(generate_sshkey_pair) || sshkey_public=$(cat $HOME/.ssh/id_ed25519.pub)
    # 3. generate cloud-init config.
    generate_vm_cloudinit
    # 4. launch vm.
    multipass launch --cloud-init $HOME/${CLOUDINIT_NAME}.yaml --name $VM_NAME
fi
# 5. make sure vm is started/running.
multipass start $VM_NAME
# 6. get ip from vm. 
vm_ip=$(get_vm_ip) && [[ -z "$vm_ip" ]] && echo "no vm ip found: exiting...." && exit 1
# 7. generate ssh config so vscode can connect to vm over ssh.
generate_sshconfig
# 8. launch vscode remote.
code --remote ssh-remote+${vm_ip}
# 9. cleanup.
rm -f $HOME/CLOUDINIT_NAME.yaml
unset VM_NAME CLOUDINIT_NAME 
