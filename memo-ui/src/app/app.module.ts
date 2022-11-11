import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { NgScrollbarModule } from 'ngx-scrollbar';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { TopicsComponent } from './pages/topics/topics.component';
import { TopicListComponent } from './components/topic-list/topic-list.component';
import { MemoListComponent } from './components/memo-list/memo-list.component';
import { TopicDetailComponent } from './pages/topic-detail/topic-detail.component';
import { TopicItemComponent } from './components/topic-list/topic-item/topic-item.component';
import { MemoItemComponent } from './components/memo-list/memo-item/memo-item.component';
import { HighlightDirective } from './directives/highlight.directive';
import { UnlessDirective } from './directives/unless.directive';
import { DropdownDirective } from './directives/dropdown.directive';
import { AuthComponent } from './pages/auth/auth.component';
import { SigninComponent } from './pages/auth/signin/signin.component';
import { SignupComponent } from './pages/auth/signup/signup.component';
import { HomeComponent } from './pages/home/home.component';
import { ShareHeaderComponent } from './share/share-header/share-header.component';
import { CurrentComponent } from './pages/auth/current/current.component';
import { AuthInterceptorService } from './services/auth-interceptor.service';
import { JobCardListComponent } from './components/job-card-list/job-card-list.component';
import { JobCardItemComponent } from './components/job-card-list/job-card-item/job-card-item.component';
import { ScrollBarComponent } from './share/scroll-bar/scroll-bar.component';
import { LightboxComponent } from './share/lightbox/lightbox.component';
import { TopicCreatingModalComponent } from './share/lightbox/topic-creating-modal/topic-creating-modal.component';
import { MemoCreatingModalComponent } from './share/lightbox/memo-creating-modal/memo-creating-modal.component';
import { MemoDetailModalComponent } from './share/lightbox/memo-detail-modal/memo-detail-modal.component';
import { CalendarComponent } from './components/calendar/calendar.component';
import { CalendarModule, DateAdapter } from 'angular-calendar';
import { adapterFactory } from 'angular-calendar/date-adapters/date-fns';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    TopicsComponent,
    TopicListComponent,
    MemoListComponent,
    TopicDetailComponent,
    TopicItemComponent,
    MemoItemComponent,
    HighlightDirective,
    UnlessDirective,
    DropdownDirective,
    AuthComponent,
    SigninComponent,
    SignupComponent,
    HomeComponent,
    ShareHeaderComponent,
    CurrentComponent,
    JobCardListComponent,
    JobCardItemComponent,
    ScrollBarComponent,
    LightboxComponent,
    TopicCreatingModalComponent,
    MemoCreatingModalComponent,
    MemoDetailModalComponent,
    CalendarComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    FontAwesomeModule,
    NgScrollbarModule,
    CalendarModule.forRoot({ provide: DateAdapter, useFactory: adapterFactory }),
  ],
  providers: [{provide: HTTP_INTERCEPTORS, useClass: AuthInterceptorService, multi: true}],
  bootstrap: [AppComponent]
})
export class AppModule { }
