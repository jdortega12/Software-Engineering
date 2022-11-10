import TestRenderer from "react-test-renderer"
import HomeScreen from "./HomeScreen"
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

test("home screen smoke test ", () => {
    const stack = createNativeStackNavigator()

    const renderer = TestRenderer.create(<HomeScreen navigation={stack}/>)
    const instance = renderer.root
})