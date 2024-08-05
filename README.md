<br />
<div id="readme-top" align="center">
  <a href="https://github.com/itsmrval/http-to-https">
    <img src="https://github.com/itsmrval/http-to-https/blob/main/logo.png" alt="Logo" width="120" height="120">
  </a>

  <h3 align="center">http-to-https</h3>

  <p align="center">
    The most simple http to https redirector
    <br />
    <br />
    <a href="https://github.com/itsmrval/http-to-https/issues">Report Bug</a>
    ·
    <a href="https://github.com/itsmrval/http-to-https/pulls">Pull request</a>
  </p>
</div>


<details>
  <summary>Table of contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">What is http-to-https ?</a>
      <ul>
        <li><a href="#built-with">Built with</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#license">License</a></li>
  </ol>
</details>



## What is http-to-https

http-to-https is a simple request redirector from http to https that log request and allow custom ports.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

This section list major frameworks/libraries used

* ![](https://img.shields.io/badge/GO-20232A?style=for-the-badge&logo=go)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



## Getting Started

Now let's see how to set up an http-to-https instance.

### Installation

1. Create directory
   ```sh
   mkdir /opt/http-to-https
   cd /opt/http-to-https
   ```
2. Download the latest release and apply permissions
   ```sh
   wget -O http-to-https https://github.com/itsmrval/http-to-https/releases/download/1.1.0/http-to-https_linux_amd64
   chmod +x http-to-https
   ```
3. Create the service on systemd
   Write the file
   ```sh
   nano /etc/systemd/system/http-to-https.service
   ```
   Complete and put the service file below:
   	```txt
   	[Unit]
	Description=http-to-https service
	After=network.target
	
	[Service]
	Type=simple
	ExecStart=/opt/http-to-https/http-to-https
	Environment="PORT=80"
	
	[Install]
	WantedBy=multi-user.target
   	```
   
6. Reload systemd and run the service !
   ```sh
   systemctl daemon-reload
   systemctl enable --now http-to-https
   ```
<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
