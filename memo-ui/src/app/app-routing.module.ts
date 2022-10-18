import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthComponent } from './auth/auth.component';
import { HomeComponent } from './home/home.component';
import { MemosComponent } from './memos/memos.component';
import { TopicsComponent } from './topics/topics.component';
import { SigninComponent } from './auth/signin/signin.component';
import { SignupComponent } from './auth/signup/signup.component';
import { MemoDetailComponent } from './memos/memo-detail/memo-detail.component';
import { TopicDetailComponent } from './topics/topic-detail/topic-detail.component';
import { AuthGuardService } from './services/auth-guard.service';

const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
  },
  {
    path: 'topic',
    component: TopicsComponent,
  },
  {
    path: 'topic/:id',
    component: TopicDetailComponent,
  },
  {
    path: 'memo',
    component: MemosComponent,
    canActivate: [AuthGuardService],
    children: [
      {
        path: ':id',
        component: MemoDetailComponent,
      }
    ]
  },
  {
    path: 'auth',
    component: AuthComponent,
    children: [
      {
        path: 'signin',
        component: SigninComponent,
      },
      {
        path: 'signup',
        component: SignupComponent,
      }
    ]
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule { }
