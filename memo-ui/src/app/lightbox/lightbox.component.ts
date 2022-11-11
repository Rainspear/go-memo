import { Component, EventEmitter, Input, OnInit, OnChanges, Output, SimpleChanges, OnDestroy } from '@angular/core';
import { Subject, Subscription } from 'rxjs';

@Component({
  selector: 'app-lightbox',
  templateUrl: './lightbox.component.html',
  styleUrls: ['./lightbox.component.scss']
})
export class LightboxComponent implements OnInit, OnChanges, OnDestroy {

  @Input() show?: boolean;
  // showSubject = new Subject<boolean>();
  @Input() toggleForm: () => void = () => {}
    // @Output() showSubject = new EventEmitter<void>();
  // @Output() onClose = new EventEmitter<boolean>();
  // subscriptionShow: Subscription;
  // @Input() onCloseApp = () => {}

  constructor() {
  }

  onClickCloseApp() {
    // this.show = !this.show;
    // this.showSubject.emit();
    this.toggleForm()
  }

  ngOnChanges(changes: any): void {
    console.log("LightboxComponent changes", changes)
    this.show = changes?.show.currentValue
  }

  ngOnInit(): void {

  }

  ngOnDestroy(): void {
    // this.subscriptionShow.unsubscribe();
  }

}
