import { Component, EventEmitter, Input, OnInit, Output, SimpleChanges } from '@angular/core';
import { Topic } from 'src/app/models/topic.model';

@Component({
  selector: 'app-topic-list',
  templateUrl: './topic-list.component.html',
  styleUrls: ['./topic-list.component.scss']
})
export class TopicListComponent implements OnInit {
  showCreatingForm: boolean = false;
  test :number = 0
  @Input() topics: Topic[] = [];
  @Output() selectedTopic = new EventEmitter<Topic>();

  onSelectTopic(topic: Topic): void {
    this.selectedTopic.emit(topic);
  }

  onToggle(showCreatingForm : boolean) {
    this.showCreatingForm = showCreatingForm;
  }

  ngOnInit(): void {
  }

  ngOnChanges(changes: SimpleChanges) {
    // console.log("changes", changes)
    // console.log("TopicListComponent ngOnChanges: ", changes)
    // this.topics = changes['topics']
  }

}
