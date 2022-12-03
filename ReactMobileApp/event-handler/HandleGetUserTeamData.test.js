import React from "react-native"
import handleGetUserTeamData from "./HandleGetUserTeamData"

// Smoke test to make sure handleGetUserTeamData() causes no
// errors.
test("handleGetUserTeamData smoke test", () => {
    handleGetUserTeamData()
})