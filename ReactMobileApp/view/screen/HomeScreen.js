// Home page

import React from "react"
import {Button, Text} from "react-native"
import HandleLogout from "../../event-handler/HandleLogout"
import TopBar from "../component/TopBar"

export default function HomeScreen() {
    return (
        <>
            <TopBar/>
            <Text>Home</Text>
            <Button title="Logout (temp)" onPress={HandleLogout}/>
        </>
    )
}