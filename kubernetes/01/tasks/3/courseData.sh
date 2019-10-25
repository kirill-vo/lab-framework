#!/bin/bash

USERNAME=user01
useradd ${USERNAME} -d /home/${USERNAME} -m -s /bin/bash
echo "${USERNAME} ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/${USERNAME}