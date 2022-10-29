/**
 * @format
 */

import 'react-native';
import React from 'react';
import App from '../App';
import TeamRequestForm from '../view/component/TeamRequestForm';
import CreateAccount from '../view/screen/CreateAccountScreen';
import HomeScreen from '../view/screen/HomeScreen';
import UpdateUserPersonalInfoScreen from '../view/screen/UpdateUserPersonalInfo';

// Note: test renderer must be required after react-native.
import renderer from 'react-test-renderer';

it('renders correctly', () => {
  renderer.create(<App />);
});

it('renders TeamRequestForm correctly with type 0', () => {
  renderer.create(<TeamRequestForm ReceiverID="0" type="0"/>);
});

it('renders TeamRequestForm correctly with type 1', () => {
  renderer.create(<TeamRequestForm ReceiverID="0" type="1" />);
});

it('renders CreateAccount correctly', () => {
  renderer.create(<CreateAccount />);
});

it('renders HomeScreen correctly', () => {
  renderer.create(<HomeScreen />);
});

it('renders UserPersonalInfo correctly', () => {
  renderer.create(<UpdateUserPersonalInfoScreen />);
});