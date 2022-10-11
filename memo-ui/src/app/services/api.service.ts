import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { Memo } from '../models/memo.model';
import { Topic } from '../models/topic.model';
import { ResponseAPI } from '../models/response.model';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'Application/json' })
}
const apiUrl = 'http://localhost:8089';

@Injectable({
  providedIn: 'root'
})
export class ApiSerivce {

  constructor(private httpClient: HttpClient) { }

  getAllMemos(): Observable<ResponseAPI<Memo[]>> {
    return this.httpClient.get<ResponseAPI<Memo[]>>(`${apiUrl}/memos`).pipe()
  }

  postMemo(data: Memo): Observable<ResponseAPI<Memo>> {
    return this.httpClient.post<ResponseAPI<Memo>>(`${apiUrl}/memos`, data).pipe()
  }

  deleteMemo(id: string): Observable<ResponseAPI<Memo>> {
    return this.httpClient.delete<ResponseAPI<Memo>>(`${apiUrl}/memos/${id}`).pipe()
  }

  getAllTopics(): Observable<ResponseAPI<Topic[]>> {
    return this.httpClient.get<ResponseAPI<Topic[]>>(`${apiUrl}/topics`).pipe()
  }
}