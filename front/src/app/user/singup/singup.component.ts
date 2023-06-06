import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Store } from '@ngrx/store';

@Component({
  selector: 'app-singup',
  templateUrl: './singup.component.html',
  styleUrls: ['./singup.component.scss', '../user.component.scss']
})
export class SingupComponent {
  singupForm!: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private http: HttpClient,
    private store: Store
  ) {}

  ngOnInit() {
    this.singupForm = this.formBuilder.group({
      name: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      password: ['', Validators.required],
      purposeOfUse: ['', Validators.required]
    });
  }

  onSubmit() {
    throw new Error('Method not implemented.');
    }
}
