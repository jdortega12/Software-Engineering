import React from "react";
import TestRenderer from "react-test-renderer";
import AskManagerRequestForm from "../component/AskManagerRequestForm";
import {Text} from "react-native"
import TeamProfile from "./TeamProfile";

test("team profile renders", () => {
    const testRenderer = TestRenderer.create(<TeamProfile id="1"/>);
});

