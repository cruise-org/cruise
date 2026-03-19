# Cruise – Container Management Client 

> Terminal UI for managing containers with style and speed.

**Cruise** is a powerful, intuitive, and fully-featured TUI (Terminal User Interface) for interacting with Containers. Built with Go and [Bubbletea](https://github.com/charmbracelet/bubbletea), it offers a visually rich, keyboard-first experience for managing containers, images, volumes, networks, logs and more — all from your terminal.

## Screenshots & Usage

<details>
  <summary>Screenshots</summary>

<img width="1827" height="1013" alt="1" src="https://github.com/user-attachments/assets/67e2903d-b754-4ea4-96f6-fb3afd3f9286" />
<img width="1827" height="1013" alt="2" src="https://github.com/user-attachments/assets/40fb2d54-e05d-4e38-96b4-07852093ee8c" />
  <img width="1827" height="1013" alt="4" src="https://github.com/user-attachments/assets/22b322db-4eef-443e-a33d-90602c238ff0" />
  <img width="1827" height="1013" alt="3" src="https://github.com/user-attachments/assets/3d4392a1-1d4c-48c3-aaa7-8f7f7e6e6576" />
<img width="1827" height="1013" alt="11" src="https://github.com/user-attachments/assets/df9fce45-d612-402d-8414-ad981ff5765c" />
<img width="1827" height="1013" alt="10" src="https://github.com/user-attachments/assets/8786024c-faf4-4f2a-ae97-d3a886b67ad0" />
<img width="1827" height="1013" alt="9" src="https://github.com/user-attachments/assets/428dc2a6-2951-4da9-ad3d-1ad4188268fa" />
<img width="1827" height="1013" alt="8" src="https://github.com/user-attachments/assets/d9dd132a-9e72-480a-b3a3-08b310e51fb4" />
<img width="1827" height="1013" alt="7" src="https://github.com/user-attachments/assets/bed515da-0c5d-4f72-b771-4f6afcfc35c4" />
<img width="1827" height="1013" alt="6" src="https://github.com/user-attachments/assets/c1999901-c2de-416c-8bc1-cfea2d368d31" />
<img width="1827" height="1013" alt="5" src="https://github.com/user-attachments/assets/fc118e59-781e-412b-b787-82f497084587" />
</details>

Once [installed](#installation). You can run the app with `cruise`.


## Description

Ever felt that CLI's for containerization tools are too lengthy or limited? Find yourself executing commands again and again for simple things? Wrote multiline commands just for a typo to ruin it? Well... Fret no more. Cruise - Is a TUI Container Management Client, fitting easily in your terminal-first dev workflow, while making repetitive work easy and interactive.

> How is _cruise_ different from existing solutions?

Existing applications are limited in what they do, they serve as mostly a monitoring service, _not_ a management service let alone a Client.

With Cruise you can:
- Manage Lifecycles of Containers, Images, Volumes, Networks.
- Have a centralized Monitoring service
- Scan images for vulnerabilities
- Get Detailed view on Docker Artifacts
- and more to come!

**NOTE**: Although cruise currently only supports docker, through some more refactor and expansion we plan to expand the scope of cruise to support multiple containerization tools and also several other QoL features.

### Tech Stack

- **Go** – High performance, robust concurrency
- **Bubbletea** – Elegant terminal UI framework
- **Charm ecosystem** – [Lipgloss](https://github.com/charmbracelet/lipgloss), [Bubbles](https://github.com/charmbracelet/bubbles), [Glamour](https://github.com/charmbracelet/glamour)
- **Docker SDK for Go** – Deep Docker integration
- **Trivy / Grype** – Vulnerability scanning
- **Viper** – Configuration management

## Future Plans

Please check out the roadmap [here](ROADMAP.md).

### Expected Growth Plan

- Support multiple popular containerization tools
- Special Support for Docker Compose.
  - Support Managing Compose Lifecycles
  - Compose Editing
  - and more...
- Support Advanced Image Manipulation.
- Advanced Monitoring Support
  - Notifications/Alerts
  - A standalone/integrated daemon service for live log/monitoring
- Registry & Image Build Configuration
  - Harbor Support?

- Support for Container Orchestration tools

**This list is just specs I can list from the top of my head, actual integration and selection will vary.**

## Installation

Refer to the installation docs [here](https://cruise-org.github.io/install).

## Contributing

Please check out [CONTRIBUTING.md](CONTRIBUTING.md) for more.


## License

MIT License – see [LICENSE](LICENSE) for details.

## Credits

Built by [Nucleo](https://github.com/NucleoFusion).

Inspiration and Advice from [SourcewareLab](https://github.com/SourcewareLab).

Special Thanks to [Hegedus Mark](https://github.com/hegedus-mark) and [Mongy](https://github.com/A-Cer23).


