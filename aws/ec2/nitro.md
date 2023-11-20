# Hypervisor 

AWS EC2 instances are virtual machines. They run on top of a hypervisor. A hypervisor is a program that creates and manages virtual machines. It is also known as a virtual machine monitor (VMM).

![](https://www.brendangregg.com/blog/images/2017/ec2-types-numbered.png)

## Xen

When EC2 was first launched, it used a hypervisor called **Xen**. Xen is an open source hypervisor that was originally developed by the University of Cambridge and is now maintained by the Linux Foundation.

## Nitro

In 2017, AWS launched a new hypervisor called **Nitro**. Nitro is a lightweight hypervisor that is designed to provide better performance and security than **Xen**.

## Bare Metal

AWS also offers bare metal instances. Bare metal instances provide direct access to the underlying hardware. They are a good choice for workloads that require access to specific hardware features that are not available in virtualized environments, such as Intel VT-x, Intel VT-d, or SR-IOV.

# Nitro System

Consists of 3 main components:

- Nitro Card
- Nitro Security Chip
- Nitro Hypervisor

## Nitro Card

The Nitro Card is a PCIe device that provides access to hardware resources. It is used to provide access to network and storage resources.

## Nitro Security Chip

The Nitro Security Chip is a dedicated hardware chip that is used to provide security services. It is used to provide secure boot, instance identity, and cryptographic services.

## Nitro Hypervisor

The Nitro Hypervisor is a lightweight hypervisor that is used to provide virtualization services. It is used to provide virtual machine management, device emulation, and memory management.

# References:

- [AWS Nitro System](https://aws.amazon.com/ec2/nitro/)
- [AWS re:Invent 2022 - Powering Amazon EC2: Deep dive on the AWS Nitro System (CMP301)](https://www.youtube.com/watch?v=jAaqfeyvvSE)