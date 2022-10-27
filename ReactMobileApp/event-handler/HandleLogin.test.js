import React from "react-native"
import handleLogin from "./HandleLogin"

// Smoke test to make sure handleLogin() causes no
// errors.
test("handleLogin smoke test", () => {
    handleLogin()
})