[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]


<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">StreamingScheduler</h3>

  <p align="center">
    Service for schedule rtmp lives
    <br />
    <a href="https://pkg.go.dev/github.com/MoraGames/StreamingScheduler"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/github_username/repo_name">View Demo</a>
    ·
    <a href="https://github.com/MoraGames/StreamingScheduler/issues">Report Bug</a>
    ·
    <a href="https://github.com/MoraGames/StreamingScheduler/issues">Request Feature</a>
  </p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

<!-- GETTING STARTED -->
## Getting Started

To download and start using the library follow these simple:

### Prerequisites

* docker
* docker-compose

### Installation

   ```sh
   git clone https://github.com/MoraGames/StreamingScheduler
   cd StreamingScheduler
   ```
---
## Usage

Let's now see how the library is used with some small examples of common use.

### Run

```sh
docker-compose up -d
```

### Use HTTPS

For use self-signed certificate run at the first run:

```sh 
sudo chmod +x .config/traefik/cert-generator.sh
sudo ./.config/traefik/cert-generator.sh streamtv.it
sudo sed -i '/^127\.0\.0\.1\s/s/$/ '"traefik.streamtv.it"'/' /etc/hosts
sudo sed -i '/^127\.0\.0\.1\s/s/$/ '"auth.streamtv.it"'/' /etc/hosts

```

---

<!-- TODO -->
## TODO

- [X] Example
- [ ] Example2



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/MoraGames/StreamingScheduler.svg?style=for-the-badge
[contributors-url]: https://github.com/MoraGames/StreamingScheduler/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/MoraGames/StreamingScheduler.svg?style=for-the-badge
[forks-url]: https://github.com/MoraGames/StreamingScheduler/network/members
[stars-shield]: https://img.shields.io/github/stars/MoraGames/StreamingScheduler.svg?style=for-the-badge
[stars-url]: https://github.com/MoraGames/StreamingScheduler/stargazers
[issues-shield]: https://img.shields.io/github/issues/MoraGames/StreamingScheduler.svg?style=for-the-badge
[issues-url]: https://github.com/MoraGames/StreamingScheduler/issues
[license-shield]: https://img.shields.io/github/license/MoraGames/StreamingScheduler.svg?style=for-the-badge
[license-url]: https://github.com/MoraGames/StreamingScheduler/blob/master/LICENSE.txt