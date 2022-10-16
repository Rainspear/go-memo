import { Directive , Input, TemplateRef, ViewContainerRef} from '@angular/core';

@Directive({
  selector: '[appUnless]'
})
export class UnlessDirective {
  /* // same as
  unless?: boolean;
  @Input() set unless(con : boolean) { this.unless = con;}
  */

  @Input() set appUnless(condition: any) {
    if (!condition) {
      this.vcRef.createEmbeddedView(this.templateRef)
    } else {
      this.vcRef.clear();
    }
  }
  constructor(private templateRef: TemplateRef<any>, private vcRef: ViewContainerRef) { }

}
