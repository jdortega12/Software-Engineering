import { StyleSheet } from "react-native"

export default StyleSheet.create({
    titleBarView: {
        backgroundColor:"#e32636",
        height: "8%",
        alignItems: 'center',
    },
    menuIconTouchArea: {
        justifyContent:"center"
    },
    searchIconTouchArea: {
        justifyContent:"center",
    },
    menuIcon: {
        color:"white",
        textAlignVertical:"center",
    },
    searchIcon: {
        color:"white",
        textAlignVertical:"center"
    },
    titleText: {
        //textAlignVertical:"center",
        //flex: 1, 
        textAlign: 'center',
        color:"white",
        fontSize:24,
    },
})