import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Topic } from 'src/app/models/topic.model';

@Component({
  selector: 'app-topic-item',
  templateUrl: './topic-item.component.html',
  styleUrls: ['./topic-item.component.scss']
})
export class TopicItemComponent implements OnInit {
  @Input() topic?: Topic;
  @Output() selectedTopic = new EventEmitter<Topic>();

  constructor() { }

  onSelectTopic() {
    this.selectedTopic.emit(this.topic);
  }

  ngOnInit(): void {
  }

}
