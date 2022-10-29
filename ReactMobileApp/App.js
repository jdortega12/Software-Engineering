import React from "react";
import { Text, View } from "react-native";
import HomeScreen from "./view/screen/HomeScreen";
import UpdateUserPersonalInfoScreen from "./view/screen/UpdateUserPersonalInfo";
import CreateAccount from "./view/screen/CreateAccountScreen"
import Login from "./view/screen/LoginScreen"
import CreateTeam from "./view/screen/CreateTeam"
//import UploadImage from "./view/screen/UploadImage"
import AskManagerRequestForm from "./view/component/AskManagerRequestForm"
import InvitePlayerRequestForm from "./view/component/InvitePlayerRequestForm"
import TeamRequestForm from "./view/component/TeamRequestForm"

// The plan for the acceptance test presentation is to 
// manually switch screens becuase we do not yet have the 
// navigation bar working.
export default function App(){
    return (
        <CreateAccount />
        //<Login />
        
        //<UpdateUserPersonalInfoScreen />
        //<HomeScreen /> // has logout button on it for demonstration

        //<CreateTeam />
        
        //<AskManagerRequestForm />
        //<InvitePlayerRequestForm />
        //<TeamRequestForm />
        //<UploadImage />
    );
}