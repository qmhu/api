//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Algorithm) DeepCopyInto(out *Algorithm) {
	*out = *in
	if in.DSP != nil {
		in, out := &in.DSP, &out.DSP
		*out = new(DSP)
		(*in).DeepCopyInto(*out)
	}
	if in.Percentile != nil {
		in, out := &in.Percentile, &out.Percentile
		*out = new(Percentile)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Algorithm.
func (in *Algorithm) DeepCopy() *Algorithm {
	if in == nil {
		return nil
	}
	out := new(Algorithm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNodePrediction) DeepCopyInto(out *ClusterNodePrediction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNodePrediction.
func (in *ClusterNodePrediction) DeepCopy() *ClusterNodePrediction {
	if in == nil {
		return nil
	}
	out := new(ClusterNodePrediction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterNodePrediction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNodePredictionList) DeepCopyInto(out *ClusterNodePredictionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterNodePrediction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNodePredictionList.
func (in *ClusterNodePredictionList) DeepCopy() *ClusterNodePredictionList {
	if in == nil {
		return nil
	}
	out := new(ClusterNodePredictionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterNodePredictionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNodePredictionSpec) DeepCopyInto(out *ClusterNodePredictionSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PredictionTemplate != nil {
		in, out := &in.PredictionTemplate, &out.PredictionTemplate
		*out = new(PredictionTemplate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNodePredictionSpec.
func (in *ClusterNodePredictionSpec) DeepCopy() *ClusterNodePredictionSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterNodePredictionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNodePredictionStatus) DeepCopyInto(out *ClusterNodePredictionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNodePredictionStatus.
func (in *ClusterNodePredictionStatus) DeepCopy() *ClusterNodePredictionStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterNodePredictionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DSP) DeepCopyInto(out *DSP) {
	*out = *in
	in.Estimators.DeepCopyInto(&out.Estimators)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DSP.
func (in *DSP) DeepCopy() *DSP {
	if in == nil {
		return nil
	}
	out := new(DSP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Estimators) DeepCopyInto(out *Estimators) {
	*out = *in
	if in.MaxValueEstimators != nil {
		in, out := &in.MaxValueEstimators, &out.MaxValueEstimators
		*out = make([]*MaxValueEstimator, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(MaxValueEstimator)
				**out = **in
			}
		}
	}
	if in.FFTEstimators != nil {
		in, out := &in.FFTEstimators, &out.FFTEstimators
		*out = make([]*FFTEstimator, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FFTEstimator)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Estimators.
func (in *Estimators) DeepCopy() *Estimators {
	if in == nil {
		return nil
	}
	out := new(Estimators)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpressionQuery) DeepCopyInto(out *ExpressionQuery) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpressionQuery.
func (in *ExpressionQuery) DeepCopy() *ExpressionQuery {
	if in == nil {
		return nil
	}
	out := new(ExpressionQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FFTEstimator) DeepCopyInto(out *FFTEstimator) {
	*out = *in
	if in.MinNumOfSpectrumItems != nil {
		in, out := &in.MinNumOfSpectrumItems, &out.MinNumOfSpectrumItems
		*out = new(int32)
		**out = **in
	}
	if in.MaxNumOfSpectrumItems != nil {
		in, out := &in.MaxNumOfSpectrumItems, &out.MaxNumOfSpectrumItems
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FFTEstimator.
func (in *FFTEstimator) DeepCopy() *FFTEstimator {
	if in == nil {
		return nil
	}
	out := new(FFTEstimator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistogramConfig) DeepCopyInto(out *HistogramConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistogramConfig.
func (in *HistogramConfig) DeepCopy() *HistogramConfig {
	if in == nil {
		return nil
	}
	out := new(HistogramConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Label) DeepCopyInto(out *Label) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Label.
func (in *Label) DeepCopy() *Label {
	if in == nil {
		return nil
	}
	out := new(Label)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaxValueEstimator) DeepCopyInto(out *MaxValueEstimator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaxValueEstimator.
func (in *MaxValueEstimator) DeepCopy() *MaxValueEstimator {
	if in == nil {
		return nil
	}
	out := new(MaxValueEstimator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricQuery) DeepCopyInto(out *MetricQuery) {
	*out = *in
	if in.QueryConditions != nil {
		in, out := &in.QueryConditions, &out.QueryConditions
		*out = make([]QueryCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricQuery.
func (in *MetricQuery) DeepCopy() *MetricQuery {
	if in == nil {
		return nil
	}
	out := new(MetricQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTimeSeries) DeepCopyInto(out *MetricTimeSeries) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		copy(*out, *in)
	}
	if in.Samples != nil {
		in, out := &in.Samples, &out.Samples
		*out = make([]Sample, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTimeSeries.
func (in *MetricTimeSeries) DeepCopy() *MetricTimeSeries {
	if in == nil {
		return nil
	}
	out := new(MetricTimeSeries)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in MetricTimeSeriesList) DeepCopyInto(out *MetricTimeSeriesList) {
	{
		in := &in
		*out = make(MetricTimeSeriesList, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(MetricTimeSeries)
				(*in).DeepCopyInto(*out)
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTimeSeriesList.
func (in MetricTimeSeriesList) DeepCopy() MetricTimeSeriesList {
	if in == nil {
		return nil
	}
	out := new(MetricTimeSeriesList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Percentile) DeepCopyInto(out *Percentile) {
	*out = *in
	out.Histogram = in.Histogram
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Percentile.
func (in *Percentile) DeepCopy() *Percentile {
	if in == nil {
		return nil
	}
	out := new(Percentile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionMetric) DeepCopyInto(out *PredictionMetric) {
	*out = *in
	if in.ResourceQuery != nil {
		in, out := &in.ResourceQuery, &out.ResourceQuery
		*out = new(corev1.ResourceName)
		**out = **in
	}
	if in.MetricQuery != nil {
		in, out := &in.MetricQuery, &out.MetricQuery
		*out = new(MetricQuery)
		(*in).DeepCopyInto(*out)
	}
	if in.ExpressionQuery != nil {
		in, out := &in.ExpressionQuery, &out.ExpressionQuery
		*out = new(ExpressionQuery)
		**out = **in
	}
	in.Algorithm.DeepCopyInto(&out.Algorithm)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionMetric.
func (in *PredictionMetric) DeepCopy() *PredictionMetric {
	if in == nil {
		return nil
	}
	out := new(PredictionMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionMetricStatus) DeepCopyInto(out *PredictionMetricStatus) {
	*out = *in
	if in.Prediction != nil {
		in, out := &in.Prediction, &out.Prediction
		*out = make([]*MetricTimeSeries, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(MetricTimeSeries)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionMetricStatus.
func (in *PredictionMetricStatus) DeepCopy() *PredictionMetricStatus {
	if in == nil {
		return nil
	}
	out := new(PredictionMetricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionTemplate) DeepCopyInto(out *PredictionTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionTemplate.
func (in *PredictionTemplate) DeepCopy() *PredictionTemplate {
	if in == nil {
		return nil
	}
	out := new(PredictionTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryCondition) DeepCopyInto(out *QueryCondition) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryCondition.
func (in *QueryCondition) DeepCopy() *QueryCondition {
	if in == nil {
		return nil
	}
	out := new(QueryCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sample) DeepCopyInto(out *Sample) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sample.
func (in *Sample) DeepCopy() *Sample {
	if in == nil {
		return nil
	}
	out := new(Sample)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesPrediction) DeepCopyInto(out *TimeSeriesPrediction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesPrediction.
func (in *TimeSeriesPrediction) DeepCopy() *TimeSeriesPrediction {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesPrediction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TimeSeriesPrediction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesPredictionList) DeepCopyInto(out *TimeSeriesPredictionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TimeSeriesPrediction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesPredictionList.
func (in *TimeSeriesPredictionList) DeepCopy() *TimeSeriesPredictionList {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesPredictionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TimeSeriesPredictionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesPredictionSpec) DeepCopyInto(out *TimeSeriesPredictionSpec) {
	*out = *in
	if in.PredictionMetrics != nil {
		in, out := &in.PredictionMetrics, &out.PredictionMetrics
		*out = make([]PredictionMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.TargetRef = in.TargetRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesPredictionSpec.
func (in *TimeSeriesPredictionSpec) DeepCopy() *TimeSeriesPredictionSpec {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesPredictionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesPredictionStatus) DeepCopyInto(out *TimeSeriesPredictionStatus) {
	*out = *in
	if in.PredictionMetrics != nil {
		in, out := &in.PredictionMetrics, &out.PredictionMetrics
		*out = make([]PredictionMetricStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesPredictionStatus.
func (in *TimeSeriesPredictionStatus) DeepCopy() *TimeSeriesPredictionStatus {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesPredictionStatus)
	in.DeepCopyInto(out)
	return out
}
