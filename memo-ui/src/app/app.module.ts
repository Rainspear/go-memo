import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { HttpClientModule } from '@angular/common/http';

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
    MemoEditingComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
