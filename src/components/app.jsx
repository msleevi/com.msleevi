// @flow

import React from 'react'
import Helmet from 'react-helmet'
import Switch from 'react-router'
import { Route } from 'react-router-dom'
import { APP_NAME } from '../constants/config'
import NavigationBar from './navigation-bar'

import ProjectsPage from './page/projects/projects'
import ContactPage from './page/contact/contact'
import HomePage from './page/home/home'
import AdminPage from './page/admin/admin'
import BlogPage from './page/blog/blog'
import CVPage from './page/cv/CV'
import Footer from './page/footer'

import NotFoundPage from './page/not-found'
import {
  HOME_PAGE_ROUTE,
  BLOG_PAGE_ROUTE,
  CONTACT_PAGE_ROUTE,
  PROJECTS_PAGE_ROUTE,
  CV_PAGE_ROUTE,
  ADMIN_PAGE_ROUTE,
} from '../server/routes/routes'


class App extends React.Component {
  render() {
   console.log("Rendering App")
   return (
      <div style={{ paddingTop: 54 }}>
        <Helmet titleTemplate={`%s | ${APP_NAME}`} defaultTitle={APP_NAME} />
        <NavigationBar />
        <Switch>
          <Route exact path={HOME_PAGE_ROUTE} render={() => <HomePage />} />
          <Route path={BLOG_PAGE_ROUTE} render={() => <BlogPage />} />
          <Route path={CONTACT_PAGE_ROUTE} render={() => <ContactPage />} />
          <Route path={PROJECTS_PAGE_ROUTE} render={() => <ProjectsPage />} />
          <Route path={CV_PAGE_ROUTE} render={() => <CVPage />} />
          <Route path={ADMIN_PAGE_ROUTE} render={() => <AdminPage />} />
          <Route component={NotFoundPage} />
        </Switch>
        <Footer />
      </div>
    )
  }
}

export { App as default }