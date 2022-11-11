import React from "react"
import UserProfileScreen from "./UserProfileScreen"
import UserProfileScreenPersonal from "./UserProfileScreenPersonal"
import UserProfileScreenNotPersonal from "./UserProfileScreenNotPersonal"
import { render } from "@testing-library/react-native"
import { END_GET_USER } from "../../../GlobalConstants"
/*
const { createServer } = require("@mocks-server/main");

const routes = [
    {
        id: "get-user",
        url: "/api/v1/get-user" + "/:username",
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
    server = createServer({
        server: {
            host: "10.0.2.2",
            port: 8080,
        }
    });

    const { loadRoutes } = server.mock.createLoaders();
    loadRoutes(routes);

    await server.start();
});
  
afterAll(async () => {
    await server.stop();
});
*/

test("Full user Profile render smoke test", () => {
    const profile = render(<UserProfileScreen username="jaluhrman" isSelf={true}/>)
})