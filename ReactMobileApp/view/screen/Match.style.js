import { StyleSheet } from "react-native";

export default StyleSheet.create({
    col: {
        flexDirection: "column",
        flex: 0.5,
        borderWidth: 3,
        marginLeft: 50,
        marginRight: 50,
    },

    col2: {
        flexDirection: "column",
    },

    col3: {
        flexDirection: "column",
        flex: 0.8,
    },

    col4: {
        flexDirection: "column",
        flex: 0.3,
        paddingLeft: 5,
        height: 55
    },  

    col5: {
        flexDirection: "column",
        flex: 0.75
    },

    row: {
        flexDirection: "row",
        flex: 0.4,
        justifyContent: "center",
    },

    row2: {
        flexDirection: "row",
        flex: 0.5,
        justifyContent: "center",
    },

    row3: {
        flexDirection: "row",
        flex: 0.3,
        justifyContent: "center",
    },

    teamcol: {
        flexDirection: "column",
        flex: 1,
        alignContent: "center",
        textAlign: "center",
    },

    textcol: {
        flexDirection: "column",
        flex: 0.5,
        alignContent: "center",
        borderWidth: 1,
        textAlign: "center",
    },

    commentrow: {
        flexDirection: "row",
        flex: .5,
        justifyContent: "center",
        borderWidth: 1,
        marginLeft: 50,
        marginRight: 50,
    },

    teamtext: {
        fontSize: 20
    },

    quartertext: {
        fontSize: 15
    },

    timetext: {
        fontSize: 25
    },

    commenttext:{
        fontSize: 15
    },

    button:{
        backgroundColor:"#D6CFC7",
        borderRadius:25,
        borderWidth: 5,
        alignItems:"center",
        justifyContent:"center",
        flex: 2.5,
      },

      buttonLike:{
        borderRadius:25,
        backgroundColor:"#e32636",
        borderWidth: 5,
        alignItems:"center",
        justifyContent:"center",
        flex:.25,
      },

    inputText:{
        fontWeight: "bold",
        color:"black",
    },

    inputView:{
        backgroundColor:"white",
        borderRadius:25, 
        borderWidth: 4,
        flexDirection: "row",
    },
});