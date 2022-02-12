## [Unreleased]
- remove all users from the channel/room when host leave the channel/room.
- sign up page nav bar bug fix - always shows logut on signup page
- ability to share link for other user to connect to the channel/room

***
## [1.1.0] - 12-02-2022
### Added
- swagger documentation
### Changed
- project struct refactored to handle multiple types of auth and rtc clients


***
## [1.0.4] - 28-10-2021
### Added
- web app modified to support agora sdk
- agora client added
- token api added to generate rtc token
### Removed
- manual websocket and webrtc web app part removed

***
## [1.0.3] - 24-10-2021
### Changed
- login issue fix

***
## [1.0.2] - 17-10-2021
### Added
- firebase admin sdk for id token verification
- custom auth middleware added for authenticated routes
- id token renew using refresh token added
- room and join room authenticated route added
- react base web client added
- need to pass id token as bearer token in auth header

***
## [1.0.1] - 14-10-2021
### Added
- configure firebase to use rest api
- sign up & login api added

### Changed
- refactored project structure

### Removed
- firebase admin sdk

***
## [1.0.0] - 14-10-2021
### Added
- firebase client added
- sign up & login api added

### Changed
- refactored project structure

### Removed
- None