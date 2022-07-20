import { createApp } from 'vue'

//import Buefy from 'buefy'
//import 'buefy/dist/buefy.css'
import App from './App.vue'
import Snap from './lib/snap'
import EventHandler from './lib/eventHandler'
import './css/style.scss'
import helper from './utils/helper'
import actionButtons from './actions/buttonActions'
//import 'bulmaSrc/utilities/_all.sass'
//import 'bulma/css/bulma.css'
//Vue.use(Buefy)

window.addEventListener('DOMContentLoaded', function () {
  const settings = helper.getServerSettings()
  let app = createApp(App);
  app.provide('settings', settings)
  app.mount('#app')
  snap()
  actionButtons.run()
})

function snap() {
  const nav = document.getElementById("app-navigation")
  const content = document.getElementById("app-content")
  console.log(content.classList.contains('no-snapper'))
  if (nav
    && !content.classList.contains('no-snapper')) {

    // App sidebar on mobile
    const snapper = new Snap({
      element: document.getElementById('app-content'),
      disable: 'right',
      maxPosition: 300, // $navigation-width
      minDragDistance: 100,
    })

    // keep track whether snapper is currently animating, and
    // prevent to call open or close while that is the case
    // to avoid duplicating events (snap.js doesn't check this)
    let animating = false
    snapper.on('animating', () => {
      // we need this because the trigger button
      // is also implicitly wired to close by snapper
      animating = true
    })
    snapper.on('animated', () => {
      animating = false
    })
    snapper.on('start', () => {
      // we need this because dragging triggers that
      animating = true
    })
    snapper.on('end', () => {
      // we need this because dragging stop triggers that
      animating = false
    })

    // These are necessary because calling open or close
    // on snapper during an animation makes it trigger an
    // unfinishable animation, which itself will continue
    // triggering animating events and cause high CPU load,
    //
    // Ref https://github.com/jakiestfu/Snap.js/issues/216
    const oldSnapperOpen = snapper.open
    const oldSnapperClose = snapper.close
    const _snapperOpen = () => {
      if (animating || snapper.state().state !== 'closed') {
        return
      }
      oldSnapperOpen('left')
    }

    const _snapperClose = () => {
      if (animating || snapper.state().state === 'closed') {
        return
      }
      oldSnapperClose()
    }

    // Needs to be deferred to properly catch in-between
    // events that snap.js is triggering after dragging.
    //
    // Skipped when running unit tests as we are not testing
    // the snap.js workarounds...

    const toggle = document.getElementById("app-navigation-toggle")
    EventHandler.add("click", toggle, function (e) {
      if (snapper.state().state !== 'left') {
        snapper.open("left")
      }
      console.log(snapper.state())
    })

    EventHandler.add("keypress", toggle, function (e) {
      if (snapper.state().state === 'left') {
        snapper.close()
      } else {
        snapper.open("left")
      }
    })
  }

}