# Iteration 1 Coverage Reports

## Golang Server App
### Statement Coverage
The main package has no test files because it's just the main function.
```text
?       jdortega12/Software-Engineering/GoServerApp     [no test files]
ok      jdortega12/Software-Engineering/GoServerApp/controller  1.462s  coverage: 82.9% of statements
ok      jdortega12/Software-Engineering/GoServerApp/model       1.641s  coverage: 88.9% of statements
```
### Branch Coverage
```main``` package (no tests):
```text
ok      jdortega12/Software-Engineering/GoServerApp     1.430s [no tests to run]

Branch coverage: 0/2
```
```controller``` package:
```text
ok      jdortega12/Software-Engineering/GoServerApp/controller  1.546s

Branch coverage: 37/52
```

```model``` package:
```text
ok      jdortega12/Software-Engineering/GoServerApp/model       1.581s

Branch coverage: 23/30
```
#### Total Branch Coverage: 60/84 (~71%)

## React Native Mobile App
```text
----------------------------------|---------|----------|---------|---------|------------------- 
File                              | % Stmts | % Branch | % Funcs | % Lines | Uncovered Line #s  
----------------------------------|---------|----------|---------|---------|------------------- 
All files                         |   65.51 |      100 |      50 |   65.51 |                    
 ReactMobileApp                   |     100 |      100 |     100 |     100 |                    
  App.js                          |     100 |      100 |     100 |     100 |                    
 ReactMobileApp/event-handler     |   95.45 |      100 |     100 |   95.45 |                    
  HandleCreateAccount.js          |     100 |      100 |     100 |     100 |                    
  HandleCreateTeam.js             |     100 |      100 |     100 |     100 |                    
  HandleLogin.js                  |     100 |      100 |     100 |     100 |                    
  HandleLogout.js                 |      75 |      100 |     100 |      75 | 11                 
  HandleUpdateUserPersonalInfo.js |     100 |      100 |     100 |     100 |                    
 ReactMobileApp/view              |       0 |        0 |       0 |       0 |                    
  Form.style.js                   |       0 |        0 |       0 |       0 |                    
 ReactMobileApp/view/component    |      40 |      100 |   28.57 |      40 |                    
  AskManagerRequestForm.js        |       0 |      100 |       0 |       0 | 7                  
  InvitePlayerRequestForm.js      |       0 |      100 |       0 |       0 | 7                  
  TeamRequestForm.js              |   41.66 |      100 |      25 |   41.66 | 13-29,42-45        
  TopBar.js                       |     100 |      100 |     100 |     100 |                    
  TopBar.style.js                 |       0 |        0 |       0 |       0 |                    
 ReactMobileApp/view/screen       |      50 |      100 |   33.33 |      50 |                    
  CreateAccountScreen.js          |      80 |      100 |      50 |      80 | 43                 
  CreateTeam.js                   |       0 |      100 |       0 |       0 | 9-34               
  HomeScreen.js                   |     100 |      100 |     100 |     100 |                    
  LoginScreen.js                  |       0 |      100 |       0 |       0 | 9-42               
  UpdateUserPersonalInfo.js       |   83.33 |      100 |      50 |   83.33 | 57                 
----------------------------------|---------|----------|---------|---------|------------------- 
                                                                                                
Test Suites: 6 passed, 6 total                                                                  
Tests:       11 passed, 11 total                                                                
Snapshots:   0 total                                                                                
Time:        5.543 s                                                                            
Ran all test suites.                                                                            
```