import { StyleSheet } from "react-native"

export default StyleSheet.create({
    titleBarView: {
        backgroundColor:"red",
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
        textAlignVertical:"center",
    },
    searchIcon: {
        textAlignVertical:"center"
    },
    titleText: {
        textAlignVertical:"center",
        fontSize:24,
    },
})