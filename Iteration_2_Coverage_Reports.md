# Iteration 2 Coverage Reports

## Golang Server App
### Statement Coverage
The main package has no test files because it's just the main function.
```text
?       jdortega12/Software-Engineering/GoServerApp     [no test files]
ok      jdortega12/Software-Engineering/GoServerApp/controller  0.076s  coverage: 84.0% of statements
ok      jdortega12/Software-Engineering/GoServerApp/model       0.057s  coverage: 94.9% of statements
```
### Branch Coverage

```main``` package (no tests):
```text
ok      jdortega12/Software-Engineering/GoServerApp     0.061s [no tests to run]

Branch coverage: 0/2
```

```controller``` package:
```text
ok      jdortega12/Software-Engineering/GoServerApp/controller  0.078s

Branch coverage: 67/92
```

```model``` package:
```text
ok      jdortega12/Software-Engineering/GoServerApp/model       0.064s

Branch coverage: 22/30
```
#### Total Branch Coverage: 89/124 (~72%)

## React Native Mobile App
```text
--------------------------------------------------------|---------|----------|---------|---------|-----------------------
File                                                    | % Stmts | % Branch | % Funcs | % Lines | Uncovered Line #s     
--------------------------------------------------------|---------|----------|---------|---------|-----------------------
All files                                               |   52.35 |       40 |   49.09 |   52.66 |                       
 ReactMobileApp                                         |     100 |      100 |     100 |     100 |                       
  App.js                                                |     100 |      100 |     100 |     100 |                       
  GlobalConstants.js                                    |     100 |      100 |     100 |     100 |                       
 ReactMobileApp/event-handler                           |   75.67 |       50 |     100 |   75.67 | 
  HandleAcceptOrDeny.js                                 |   22.22 |       50 |     100 |   22.22 | 7-20
  HandleCreateAccount.js                                |     100 |      100 |     100 |     100 | 
  HandleCreateTeam.js                                   |     100 |      100 |     100 |     100 | 
  HandleGetUserProfile.js                               |   83.33 |      100 |     100 |   83.33 | 7
  HandleLogin.js                                        |     100 |      100 |     100 |     100 | 
  HandleLogout.js                                       |      75 |      100 |     100 |      75 | 11
  HandleUpdateUserPersonalInfo.js                       |     100 |      100 |     100 |     100 | 
 ReactMobileApp/event-handler/request_manager_promotion |   63.63 |      100 |      50 |   63.63 | 
  HandleGetPromotionToManagerRequests.js                |   83.33 |      100 |     100 |   83.33 | 9
  HandleRequestToBeManager.js                           |      40 |      100 |   33.33 |      40 | 16-19
 ReactMobileApp/view                                    |       0 |        0 |       0 |       0 | 
  Form.style.js                                         |       0 |        0 |       0 |       0 | 
  Notifications.style.js                                |       0 |        0 |       0 |       0 | 
  Profile.style.js                                      |       0 |        0 |       0 |       0 | 
  Team.style.js                                         |       0 |        0 |       0 |       0 | 
 ReactMobileApp/view/component                          |   53.33 |      100 |   57.14 |   53.33 | 
  AskManagerRequestForm.js                              |     100 |      100 |     100 |     100 | 
  InvitePlayerRequestForm.js                            |     100 |      100 |     100 |     100 | 
  NavBar.js                                             |       0 |        0 |       0 |       0 | 
  NavBar.style.js                                       |       0 |        0 |       0 |       0 | 
  TeamRequestForm.js                                    |   41.66 |      100 |      25 |   41.66 | 13-28,41-44
  TopBar.js                                             |     100 |      100 |     100 |     100 | 
  TopBar.style.js                                       |       0 |        0 |       0 |       0 | 
 ReactMobileApp/view/screen                             |      50 |    16.66 |   30.76 |      50 | 
  AcceptOrDeny.js                                       |      80 |      100 |      50 |      80 | 41
  CreateAccountScreen.js                                |      80 |      100 |      50 |      80 | 43
  CreateTeam.js                                         |      75 |      100 |      50 |      75 | 34                    
  HomeScreen.js                                         |   11.11 |      100 |   11.11 |   11.11 | 17-45
  LoginScreen.js                                        |      75 |      100 |      50 |      75 | 34
  TeamProfile.js                                        |   34.78 |    16.66 |   28.57 |   34.78 | 21,41-49,61-75,96-130
  UpdateUserPersonalInfo.js                             |   83.33 |      100 |      50 |   83.33 | 57
 ReactMobileApp/view/screen/admin_notifications         |   28.57 |       50 |   66.66 |   30.76 | 
  AdminNotificationsScreen.js                           |   28.57 |       50 |   66.66 |   30.76 | 17,36-62
 ReactMobileApp/view/screen/user_profile                |   23.33 |     37.5 |   42.85 |   23.33 | 
  UserProfileScreen.js                                  |      25 |     37.5 |      60 |      25 | 26,52-87
  UserProfileScreenNotPersonal.js                       |       0 |      100 |       0 |       0 | 5
  UserProfileScreenPersonal.js                          |       0 |      100 |       0 |       0 | 5
--------------------------------------------------------|---------|----------|---------|---------|-----------------------

Test Suites: 20 passed, 20 total
Tests:       25 passed, 25 total
Snapshots:   0 total
Time:        5.897 s
Ran all test suites.                                                                  
```