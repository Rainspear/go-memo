import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { NgScrollbarModule } from 'ngx-scrollbar';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { TopicsComponent } from './topics/topics.component';
import { TopicListComponent } from './topics/topic-list/topic-list.component';
import { MemosComponent } from './memos/memos.component';
import { MemoListComponent } from './memos/memo-list/memo-list.component';
import { MemoDetailComponent } from './memos/memo-detail/memo-detail.component';
import { TopicDetailComponent } from './topics/topic-detail/topic-detail.component';
import { TopicItemComponent } from './topics/topic-list/topic-item/topic-item.component';
import { MemoEditingComponent } from './memos/memo-editing/memo-editing.component';
import { MemoItemComponent } from './memos/memo-list/memo-item/memo-item.component';
import { HighlightDirective } from './directives/highlight.directive';
import { UnlessDirective } from './directives/unless.directive';
import { DropdownDirective } from './directives/dropdown.directive';
import { AuthComponent } from './auth/auth.component';
import { SigninComponent } from './auth/signin/signin.component';
import { SignupComponent } from './auth/signup/signup.component';
import { HomeComponent } from './home/home.component';
import { ShareHeaderComponent } from './share/share-header/share-header.component';
import { CurrentComponent } from './auth/current/current.component';
import { AuthInterceptorService } from './services/auth-interceptor.service';
import { JobCardComponent } from './jobs/job-card/job-card.component';
import { JobCardListComponent } from './jobs/job-card/job-card-list/job-card-list.component';
import { JobCardItemComponent } from './jobs/job-card/job-card-list/job-card-item/job-card-item.component';
import { ScrollBarComponent } from './share/scroll-bar/scroll-bar.component';
import { LightboxComponent } from './lightbox/lightbox.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    TopicsComponent,
    TopicListComponent,
    MemosComponent,
    MemoListComponent,
    MemoDetailComponent,
    TopicDetailComponent,
    TopicItemComponent,
    MemoEditingComponent,
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
    JobCardComponent,
    JobCardListComponent,
    JobCardItemComponent,
    ScrollBarComponent,
    LightboxComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    FontAwesomeModule,
    NgScrollbarModule,
  ],
  providers: [{provide: HTTP_INTERCEPTORS, useClass: AuthInterceptorService, multi: true}],
  bootstrap: [AppComponent]
})
export class AppModule { }
