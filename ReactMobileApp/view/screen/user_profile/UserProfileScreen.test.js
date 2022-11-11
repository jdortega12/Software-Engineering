import React from "react"
import UserProfileScreen from "./UserProfileScreen"
import UserProfileScreenPersonal from "./UserProfileScreenPersonal"
import UserProfileScreenNotPersonal from "./UserProfileScreenNotPersonal"
import TestRenderer from "react-test-renderer"
import { END_GET_USER } from "../../../GlobalConstants"

const { createServer } = require("@mocks-server/main");

const routes = [
    {
        id: "get-user",
        url: END_GET_USER + "/jaluhrman",
        method: "GET",
        variants: [
            {
                id: "success",
                type: "json",
                options: {
                    status: 302,
                    body: [
                        {
                            user: {
                                username: "jaluhrman",
                                email: "jaluhrman@whatever.com",
                                role: "player",
                                position: "none",
                                photo: "",
                            },
                            personal_info: {
                                firstname: "Joe",
                                lastname: "Luhrman",
                                height: 50,
                                weight: 50,
                            },
                            team_name: "badgers",
                        }
                    ] 
                }
            }
        ]
    }
]

let server

beforeAll(async () => {
    server = createServer();

    const { loadRoutes } = server.mock.createLoaders();
    loadRoutes(routes);

    await server.start();
});
  
afterAll(async () => {
    await server.stop();
});

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
    const renderer = TestRenderer.create(<UserProfileScreen username="jaluhrman" isSelf={true}/>)
    const instance = renderer.root
})