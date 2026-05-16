# WorldTime SDK feature factory

from feature.base_feature import WorldTimeBaseFeature
from feature.test_feature import WorldTimeTestFeature


def _make_feature(name):
    features = {
        "base": lambda: WorldTimeBaseFeature(),
        "test": lambda: WorldTimeTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
