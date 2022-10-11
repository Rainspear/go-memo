import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable,of} from 'rxjs';
import { Memo } from '../models/memo.model';
import { Topic } from '../models/topic.model';

const httpOptions ={
  headers:new HttpHeaders({'Content-Type':'Application/json'})
}
const apiUrl = 'http://localhost:8089';

@Injectable({
  providedIn: 'root'
})
export class ApiSerivce {

  constructor(private httpClient:HttpClient) { }

  getAllMemos():Observable<Memo[]>{
    return this.httpClient.get<Memo[]>(`${apiUrl}/memos`).pipe(
    )
  }

  getAllTopics():Observable<Topic[]>{
    return this.httpClient.get<Topic[]>(`${apiUrl}/topics`).pipe(
    )
  }
}