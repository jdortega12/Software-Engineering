import { StyleSheet } from "react-native";

export default StyleSheet.create({
    container: {
        alignItems: 'center',
        justifyContent: 'center',
        flex: 0
    },

    nameText: {
        fontSize: 35,
        fontWeight: 'bold',
    },

    playerText: {
        fontSize: 20,
        fontWieght: 'bold'
    },

    locationText: {
        fontSize: 20,
    },

    row: {
        justifyContent: 'center',
        flexDirection: 'row',
        alignItems: 'center',
        marginTop: 10,
    },

    playerRow: {
        justifyContent: 'flex-start',
        flexDirection: 'row',
        alignItems: 'center',
        marginTop: 10,
        left: 75
    },

    col: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
    },
});