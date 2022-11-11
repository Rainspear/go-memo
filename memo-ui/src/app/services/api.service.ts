import { Injectable, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { Memo } from '../models/memo.model';
import { FilterParamsTopic, Topic } from '../models/topic.model';
import { ResponseAPI } from '../models/response.model';
import { ParamsCreateUser, ParamsLoginUser, User } from '../models/user.model';
import { AuthComponent } from '../auth/auth.component';
import { Schedule } from '../models/schedule.model';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'Application/json' })
}
const apiUrl = 'http://localhost:8089';

@Injectable({
  providedIn: 'root'
})
export class ApiService implements OnInit {

  constructor(private httpClient: HttpClient) { }

  ngOnInit() {
  }

  getAllMemos(): Observable<ResponseAPI<Memo[]>> {
    return this.httpClient.get<ResponseAPI<Memo[]>>(`${apiUrl}/memos`).pipe()
  }

  getSingleMemo(id: string): Observable<ResponseAPI<Memo>> {
    return this.httpClient.get<ResponseAPI<Memo>>(`${apiUrl}/memos/${id}`).pipe()
  }

  postMemo(data: Memo): Observable<ResponseAPI<Memo>> {
    return this.httpClient.post<ResponseAPI<Memo>>(`${apiUrl}/memos`, data, httpOptions).pipe()
  }

  deleteMemo(id: string): Observable<ResponseAPI<Memo>> {
    return this.httpClient.delete<ResponseAPI<Memo>>(`${apiUrl}/memos/${id}`).pipe()
  }

  getAllTopics(): Observable<ResponseAPI<Topic[]>> {
    return this.httpClient.get<ResponseAPI<Topic[]>>(`${apiUrl}/topics`).pipe()
  }

  deleteTopic(id: string): Observable<ResponseAPI<Topic>> {
    return this.httpClient.delete<ResponseAPI<Topic>>(`${apiUrl}/topics/${id}`).pipe()
  }

  getSingleTopic(id: string, filter?: FilterParamsTopic): Observable<ResponseAPI<Topic[]>> {
    // URLSearchParams
    return this.httpClient.get<ResponseAPI<Topic[]>>(`${apiUrl}/topics/${id}`, { params: { ...filter } }).pipe()

  }

  createUser(user: ParamsCreateUser): Observable<ResponseAPI<User>> {
    return this.httpClient.post<ResponseAPI<User>>(`${apiUrl}/signup`, user).pipe()
  }

  currentUser(): Observable<ResponseAPI<User>> {
    return this.httpClient.get<ResponseAPI<User>>(`${apiUrl}/current-user`).pipe()
  }

  logOutUser(): Observable<ResponseAPI<string>> {
    return this.httpClient.post<ResponseAPI<string>>(`${apiUrl}/signout`, {}).pipe()
  }

  login(user: ParamsLoginUser): Observable<ResponseAPI<User>> {
    return this.httpClient.post<ResponseAPI<User>>(`${apiUrl}/signin`, user).pipe()
  }

  getScheduleByTopicId(topicId: string): Observable<ResponseAPI<Schedule[]>> {
    return this.httpClient.get<ResponseAPI<Schedule[]>>(`${apiUrl}/schedules`, { params: { topic_id: topicId } }).pipe()
  }
}