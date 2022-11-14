import { StyleSheet } from "react-native";

export default StyleSheet.create({
    playerCol: {
        justifyContent: 'center',
        flexDirection: 'column',
        alignItems: 'center',
        marginTop: 10,
        left: 15
    },

    bracketBox: {
        flexDirection: 'column',
        margin: 5
    },

    playoffContainer: {
        flexDirection: 'column',
        alignItems: 'center',
        flex: 0
    },

    container: {
        justifyContent: 'center',
        alignItems: 'center'
    },

    row: {
        flexDirection: 'row', 
        width: 200,
        borderWidth: 5,
        margin: 15,
    },

    // I know the name is bad im tired
    row2: {
        flexDirection: 'row', 
        margin: 15,
    },  

    roundText: {
        fontSize: 25, 
        fontWeight: 'bold'
    },

    button: {
        backgroundColor:"#e32636",
        borderRadius:25,
        height:50,
        alignItems:"center",
        justifyContent:"center",
        marginBottom:10,
        marginLeft: 15,
        width: 50
    }
})