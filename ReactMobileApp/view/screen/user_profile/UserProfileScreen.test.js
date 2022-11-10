import React from "react"
import UserProfileScreen from "./UserProfileScreen"
import UserProfileScreenPersonal from "./UserProfileScreenPersonal"
import UserProfileScreenNotPersonal from "./UserProfileScreenNotPersonal"
import TestRenderer from "react-test-renderer"
//import { createServer } from "miragejs"
//import { END_GET_USER } from "../../../GlobalConstants"

/*
let server 

beforeEach(() => {
    server = createServer()
})
  
afterEach(() => {
    server.shutdown()
})*/

test("User Profile Render Smoke Test", () => {
    <UserProfileScreen username="bingus" isSelf={true}/>
})

test("Personal Profile Render Smoke Test", () => {
    eh = UserProfileScreenPersonal("test")
})

test("Generic Profile Render Smoke Test", () => {
    meh = UserProfileScreenNotPersonal("test")
})

test("Full user Profile render smoke test", () => {
    const testRenderer = TestRenderer.create(<UserProfileScreen username="jaluhrman" isSelf={true}/>)
    const testInstance = testRenderer.root
})