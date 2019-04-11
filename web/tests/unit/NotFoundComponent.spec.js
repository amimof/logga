import { shallowMount } from '@vue/test-utils'
import NotFoundComponent from '@/components/NotFoundComponent.vue'

describe('NotFoundComponent.vue', () => {
  it('renders props.msg when passed', () => {
    const msg = 'new message'
    const wrapper = shallowMount(NotFoundComponent, {
      propsData: { msg }
    })
    expect(wrapper.text()).toMatch(msg)
  })
  it('renders default msg when props.msg not passed', () => {
    const defaultMsg = 'Ooops! Page not found'
    const wrapper = shallowMount(NotFoundComponent, {})
    expect(wrapper.text()).toMatch(defaultMsg)
  })
})
