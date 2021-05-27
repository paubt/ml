#   Challenger Space Shuttle O-Ring Data
    
###  Reference:

    David Draper (draper@math.ucla.edu)
    University of California, Los Angeles
    Donor: David Draper (draper@math.ucla.edu)
    Date: 5 August 1993

### Source:

[click here](https://archive.ics.uci.edu/ml/datasets/Challenger+USA+Space+Shuttle+O-Ring)

###  Discussion:

    There are two databases: (both use the same set of 5 attributes)
     1. Primary o-ring erosion and/or blowby
     2. Primary o-ring erosion only
    The two databases are identical except for the 2nd attribute of the
    21st instance (confirmed by David Draper on 8/5/93).
    
    Edited from (Draper, 1993):
    The motivation for collecting this database was the explosion of the
    USA Space Shuttle Challenger on 28 January, 1986.  An investigation
    ensued into the reliability of the shuttle's propulsion system.  The
    explosion was eventually traced to the failure of one of the three field
    joints on one of the two solid booster rockets.  Each of these six field
    joints includes two O-rings, designated as primary and secondary, which
    fail when phenomena called erosion and blowby both occur.
    The night before the launch a decision had to be made regarding
    launch safety.  The discussion among engineers and managers leading to
    this decision included concern that the probability of failure of the
    O-rings depended on the temperature t at launch, which was forecase to
    be 31 degrees F. There are strong engineering reasons based on the
    composition of O-rings to support the judgment that failure
    probability may rise monotonically as temperature drops.  One other
    variable, the pressure s at which safety testing for field join leaks
    was performed, was available, but its relevance to the failure process
    was unclear.
    Draper's paper includes a menacing figure graphing the number of field
    joints experiencing stress vs. liftoff temperature for the 23 shuttle
    flights previous to the Challenger disaster.  No previous liftoff
    temperature was under 53 degrees F.  Although tremendous extrapolation
    must be done from the given data to assess risk at 31 degrees F, it
    is obvious even to the layman "to foresee the unacceptably high risk
    created by launching at 31 degrees F."  For more information, see
    Draper (1993) or the other previous analyses.
    The task is to predict the number of O-rings that will experience
    thermal distress for a given flight when the launch temperature is
    below freezing

### Task:
    Number of Attributes: 5
     1. Number of O-rings at risk on a given flight
     2. Number experiencing thermal distress
     3. Launch temperature (degrees F)
     4. Leak-check pressure (psi)
     5. Temporal order of flight
    
    Attribute Information: all values are positive integers
#### created:
    pau
    27.05.2021

#### usage:
    linear regression 
