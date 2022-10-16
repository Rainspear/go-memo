import { Injectable, EventEmitter } from '@angular/core';
import { Topic } from '../models/topic.model';
@Injectable({
  providedIn: 'root'
})
export class TopicSelectingService {
  topic?: Topic
  selectedTopic = new EventEmitter<Topic>();

  constructor() { 

  }

  onGetTopic(topic: Topic) {
    this.topic = topic
    this.selectedTopic.emit(topic);
  }
}
