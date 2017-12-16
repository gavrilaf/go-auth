import unittest
import spawn_api as spawn
import uuid


class TestProfile(unittest.TestCase):

    def setUp(self):
        self.client = spawn.TEST_CLEINT
        self.endpoint = 'http://localhost:8080'
        self.api = spawn.SpawnApi(self.endpoint, self.client)

    @staticmethod
    def get_name():
        return str(uuid.uuid4()) + "@spawn.com"

    def testGetState(self):
        username = self.get_name()
        device = spawn.Device("test-device-1", "test-device-1-name")
        password = "password"

        err = self.api.sign_up(username, password, device, "ru", "es")
        self.assertIsNone(err)

        is_error, state = self.api.get_state()
        self.assertFalse(is_error, "Error is {}".format(state))

        self.assertEqual("es", state["lang"])
        self.assertEqual("ru", state["locale"])
        self.assertEqual(self.api.permissions, state["permissions"])

    def testLogout(self):
        username = self.get_name()
        device = spawn.Device("test-device-1", "test-device-1-name")
        password = "password"

        err = self.api.sign_up(username, password, device, "ru", "es")
        self.assertIsNone(err)

        is_error, state = self.api.get_state()
        self.assertFalse(is_error, "Error is {}".format(state))

        err = self.api.logout()
        self.assertIsNone(err)

        # should be error
        is_error, err = self.api.get_state()
        self.assertTrue(is_error)
        self.assertEqual(err["scope"], "auth")
        self.assertEqual(err["reason"], "session-not-found")

    def testGetDevices(self):
        username = self.get_name()
        device = spawn.Device("test-device-1", "test-device-1-name")
        password = "password"

        err = self.api.sign_up(username, password, device, "it", "es")
        self.assertIsNone(err)

        err = self.api.sign_in(username, password, spawn.Device("test-device-1-new", "test-device-1-name-new"), "ru",
                               "en")
        self.assertIsNone(err)

        is_error, devices = self.api.get_devices()
        self.assertFalse(is_error)

        devices = devices["devices"]
        self.assertEqual(2, len(devices))

        first = [x for x in devices if x["device_id"] == "test-device-1"]
        self.assertEqual(1, len(first))

        self.assertEqual("test-device-1-name", first[0]["device_name"])
        self.assertEqual(True, first[0]["is_confirmed"])
        self.assertEqual(False, first[0]["is_current"])
        self.assertEqual("it", first[0]["locale"])
        self.assertEqual("es", first[0]["lang"])
        self.assertIsNotNone(first[0]["login_ip"])
        self.assertIsNotNone(first[0]["login_region"])
        self.assertIsNotNone(first[0]["login_time"])
        self.assertIsNotNone(first[0]["user_agent"])

        current = [x for x in devices if x["device_id"] == "test-device-1-new"]
        self.assertEqual(1, len(current))

        self.assertEqual("test-device-1-name-new", current[0]["device_name"])
        self.assertEqual(False, current[0]["is_confirmed"])
        self.assertEqual(True, current[0]["is_current"])
        self.assertEqual("ru", current[0]["locale"])
        self.assertEqual("en", current[0]["lang"])
        self.assertIsNotNone(current[0]["login_ip"])
        self.assertIsNotNone(current[0]["login_region"])
        self.assertIsNotNone(current[0]["login_time"])
        self.assertIsNotNone(current[0]["user_agent"])

if __name__ == '__main__':
    unittest.main()
