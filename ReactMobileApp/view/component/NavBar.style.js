import { StyleSheet } from "react-native"

export default StyleSheet.create({
    titleBarView: {
        backgroundColor:"#e32636",
        height: "8%",
        flexDirection:"row",
        justifyContent:"center",
        alignItems: 'center',
    },
    titleText: {
        textAlignVertical:"center",
        color:"white",
        fontSize:24,
    },
    centeredText: {
        flex: 1, 
        alignItems: 'center', 
        justifyContent: 'center', 
    },
    homeText: {
        backgroundColor: "#e32636",
        color:"white",
        fontSize:24,
    },
    iconText: {
        color:"white",
        fontSize:20,
        textAlign: "center",
    },
    icon: {
        color: "white",
        fontSize: 30,
    }

})