// Title bar shared across all pages

import React from "react"
import {Text, View, TouchableOpacity} from "react-native"
import TopBarStyle from "./TopBar.style"


// Default title bar shared across all pages. Couldn't get the 
// actual icons working so for now menu and search are just text.
export default function topBar() {
    return (
        <View style={TopBarStyle.titleBarView}>    
            <Text style={TopBarStyle.titleText}>Fake Football League</Text>
        </View>

    )    
}