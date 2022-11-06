import { StyleSheet } from "react-native"

export default StyleSheet.create({
    headerSectionView: {
        flex: 0,
        paddingHorizontal: 5,
        margin: 10,
        borderColor: "grey",
        borderBottomWidth:3,
        borderRadius:15,
        backgroundColor: 'white',
        alignItems: "flex-start",
    },
    photoSectionView: {
        flex: 1.10,
        margin: 20,
        backgroundColor: 'white',
        justifyContent:'space-between',
        flexDirection:'row'
    },
    titleSectionView: {
        flex: 0,
        borderBottomWidth: 2,
        borderTopWidth: 2,
        margin: 10,
        borderColor: "grey",
        backgroundColor: 'white',
        paddingHorizontal: 5,
        justifyContent:"space-between",
        flexDirection:'row'
    },
    interactionsSectionView: {
        borderWidth:3,
        borderRadius:15,
        flex: 1.5,
        margin: 10,
        backgroundColor: "white",
        alignItems: "flex-start",
        borderColor:'grey'
    },
    emptySpaceView: {
        flex: 4,
        margin: 10,
    },
    headerText: {
        color: "black",
        fontWeight:'bold',
        fontSize: 25,
        margin:5,
    },
    titleText: {
        fontWeight:'bold',
        fontSize: 20,
        color: "black",
        textAlignVertical:'top',
        flexDirection:'row'
    },
    subtitleText: {
        color:'grey',
        fontWeight: '600',
        fontSize: 16,
        textAlignVertical:'top'
    },
    proilePicImage: {
        alignSelf:'center',
        width: 100,
        height: 100,
        borderRadius: 60/2,
    },
    likesText: {
        textAlignVertical:'center',
        textAlign:'center'
    },
    interactionArea: {
        backgroundColor: "#e32636",
        justifyContent:'center',
        flex: 1,
        width: 100,
        margin: 10,
        borderRadius: 10,
        borderColor:"black",
        borderWidth:2,
    },
    interactionText: {
        textAlign: 'center',
        textAlignVertical: 'center',
        color:'black',
        fontWeight: 'bold',
        fontSize: 11,
        margin:1,
    }
})