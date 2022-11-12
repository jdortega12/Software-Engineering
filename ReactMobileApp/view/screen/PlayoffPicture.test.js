import React from "react";
import TestRenderer from "react-test-renderer";
import PlayoffPicture from "./PlayoffPicture"

test("playoff picture renders", () => {
    const testRenderer = TestRenderer.create(<PlayoffPicture />);
});