import { StyleSheet } from "react-native"

export default StyleSheet.create({
    titleBarView: {
        backgroundColor:"#e32636",
        height: "8%",
        flexDirection:"row",
        justifyContent:"space-between",
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
        textAlignVertical:"center",
        color:"white",
        fontSize:24,
    },
})