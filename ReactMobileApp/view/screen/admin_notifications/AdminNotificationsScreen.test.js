import React from "react"
import TestRenderer from "react-test-renderer"
import AdminNotifcationsScreen from "./AdminNotificationsScreen"

test("admind notifications screen smoke test", () => {
    const testRenderer = TestRenderer.create(<AdminNotifcationsScreen/>)
    const testInstance = testRenderer.root
})