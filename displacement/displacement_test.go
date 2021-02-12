package main

import "testing"

func TestDisplacement(t *testing.T) {

	/*When time is zero and everything is zero*/
	displacementFun := GenDisplaceFn(0.0, 0.0, 0.0)
	displacement := displacementFun(0)
	if displacement != 0 {
		t.Errorf("Displacement function not returning 52 returning %f", displacement)
	}

	/* Using displacement values of (10, 2, 1) and time 0 */
	displacementFun = GenDisplaceFn(10, 2, 1)
	displacement = displacementFun(0)
	if displacement != 1 {
		t.Errorf("Displacement function not returning 52 returning %f", displacement)
	}

	/* Using same displacement formula but with a time 3 */
	displacement = displacementFun(3)
	if displacement != 52 {
		t.Errorf("Displacement function not returning 52 returning %f", displacement)
	}

	/* Using same displacement formula but with a time 5 */
	displacement = displacementFun(5)
	if displacement != 136 {
		t.Errorf("Displacement function not returning 52 returning %f", displacement)
	}
}
