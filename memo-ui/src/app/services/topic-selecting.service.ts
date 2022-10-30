import { Injectable, EventEmitter } from '@angular/core';
import { Subject } from 'rxjs';
import { IFilterTopic, Topic } from '../models/topic.model';
@Injectable({
  providedIn: 'root'
})
export class TopicSelectingService {
  topic?: Topic
  // selectedTopic = new EventEmitter<Topic>();
  selectedTopic = new Subject<Topic>()
  selectedFilter = new Subject<IFilterTopic>()

  constructor() { 

  }

  // onGetTopic(topic: Topic) {
  //   this.topic = topic
  //   this.selectedTopic.emit(topic);
  // }
}
