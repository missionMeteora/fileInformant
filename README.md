# File Informant [![GoDoc](https://godoc.org/github.com/missionMeteora/fileInformant?status.svg)](https://godoc.org/github.com/missionMeteora/fileInformant) ![Status](https://img.shields.io/badge/status-beta-yellow.svg)

File Informant is a stand-alone service which is used to actively monitor files. Through the config,
files to monitor can be set at independent intervals. If the file has not changed in size, or has
not been modified within the specified interval - all subscribers (declared in the config) will be notified.
The methods of notification are SMS (Using Twilio) and Email (Using Mandril). 

*This service will work with GCE, some email services do not due to GCE blocking the default SMTP port.*
