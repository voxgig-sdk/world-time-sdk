
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { WorldTimeSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await WorldTimeSDK.test()
    equal(null !== testsdk, true)
  })

})
