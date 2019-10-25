#!/bin/bash

USERNAME=user01
useradd ${USERNAME} -d /home/${USERNAME} -m
echo "${USERNAME} ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/${USERNAME}